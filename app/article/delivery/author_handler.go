package delivery

import (
	"go_clean_arch_test/app/article/delivery/request"
	response "go_clean_arch_test/app/article/delivery/response"
	"go_clean_arch_test/app/article/usecase"
	loginUsecase "go_clean_arch_test/app/article/usecase/auth"
	"go_clean_arch_test/app/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

const NO_AUTHORS = "no authors"
const EXISTS = "exists"

// AuthorHandler interface
type AuthorHandler interface {
	GetAllAuthor(ctx *gin.Context)
	InputAuthor(ctx *gin.Context)
	UpdateAuthor(ctx *gin.Context)
	DeleteAuthor(ctx *gin.Context)
}

type authorHandler struct {
	articleusecase usecase.ArticleUsecase
	authorUsecase  usecase.AuthorUsecase
	loginUsecase   loginUsecase.LoginUsecase
}

// NewAuthorHandler constructor
func NewAuthorHandler(articleusecase usecase.ArticleUsecase, authorUsecase usecase.AuthorUsecase, loginUsecase loginUsecase.LoginUsecase) AuthorHandler {
	return &authorHandler{articleusecase: articleusecase, authorUsecase: authorUsecase, loginUsecase: loginUsecase}
}

// 全件取得
func (authorHandler *authorHandler) GetAllAuthor(ctx *gin.Context) {

	user, err := authorHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	authors, expPool, err := authorHandler.authorUsecase.GetByUser(user.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	if len(authors) < 0 || authors == nil {
		ctx.JSON(http.StatusOK, NewH(NO_AUTHORS, authors))
		return
	}

	var authorResponseSlice []response.Author
	for _, author := range authors {
		authorResponse := response.NewAuthor(author.GetId(), author.GetName(), author.GetUserId(), author.GetUpdatedAt(), author.GetCreatedAt())
		authorResponseSlice = append(authorResponseSlice, *authorResponse)
	}

	expPoolResponse := response.NewExpPool(expPool.GetId(), expPool.GetUserId(), expPool.GetExp(), expPool.GetLv(), expPool.GetUpdatedAt(), expPool.GetCreatedAt())

	authorGetByAllResponse := response.NewAuthorGetByAllResponse(authorResponseSlice, *expPoolResponse)

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), authorGetByAllResponse))
}

// 新規登録
func (authorHandler *authorHandler) InputAuthor(ctx *gin.Context) {

	author := request.Author{}
	err := ctx.Bind(&author)
	if err != nil {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), author))
		return
	}

	user, err := authorHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	authorModel, _ := domain.NewAuthor(author.Id, author.Name, user.Id, author.UpdatedAt, author.CreatedAt)

	// カテゴリー検索(カテゴリー名で)
	authorByName, err := authorHandler.authorUsecase.GetByName(authorModel.GetName(), authorModel.GetUserId())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	if authorByName.GetId() != 0 {
		ctx.JSON(http.StatusOK, NewH(EXISTS, author))
		return

	}

	// 新規登録
	authorByName, err = authorHandler.authorUsecase.Input(ctx, authorModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), author))
}

// 更新
func (authorHandler *authorHandler) UpdateAuthor(ctx *gin.Context) {

	author := request.Author{}
	err := ctx.Bind(&author)
	if err != nil {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), author))
		return
	}

	user, err := authorHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	authorModel, _ := domain.NewAuthor(author.Id, author.Name, user.Id, author.UpdatedAt, author.CreatedAt)

	authorByName, err := authorHandler.authorUsecase.GetByName(authorModel.GetName(), authorModel.GetUserId())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	if authorByName.GetId() != 0 {
		ctx.JSON(http.StatusOK, NewH(EXISTS, author))
		return
	}

	err = authorHandler.authorUsecase.Update(ctx, authorModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), author))
}

// 削除(id指定)
func (authorHandler *authorHandler) DeleteAuthor(ctx *gin.Context) {
	author := request.Author{}
	err := ctx.Bind(&author)
	if err != nil {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), author))
		return
	}

	// ログインユーザーID且つ、記事ID データが存在しなければ、302で返す
	user, err := authorHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	authorModel, _ := domain.NewAuthor(author.Id, author.Name, user.Id, author.UpdatedAt, author.CreatedAt)

	articleByIdAndUserId, err := authorHandler.authorUsecase.GetByAuthorIdAndUserId(authorModel.GetId(), user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	if articleByIdAndUserId.GetId() == 0 {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), author))
		return
	}

	// エラーじゃなければ(削除件数1以上)、紐づくarticle削除
	err = authorHandler.articleusecase.DeleteByAuthor(ctx, authorModel, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), authorModel.GetId()))
}
