package response

/*
* API要求に対するレスポンス表現
 */
type Result struct {
	Status      int
	Message     string
	ErrorReason string
	Result      interface{}
}

/* 汎用レスポンスキー */
var (
	// OK(200)
	SUCCESS_GENERAL          = Result{Status: 200, Message: "要求処理に成功しました"}
	SUCCESS_LOGIN            = Result{Status: 200, Message: "ログインに成功しました"}
	SUCCESS_ACCOUNT_REGISTER = Result{Status: 200, Message: "アカウントの作成に成功しました"}
	SUCCESS_ACCOUNT_CHANGE   = Result{Status: 200, Message: "アカウントの変更に成功しました"}
	// BadRequest(400)
	ERROR_REQUEST_GENERAL           = Result{Status: 400, Message: "入力内容を確認してください"}
	ERROR_REQUEST_EMAIL             = Result{Status: 400, Message: "正しいメールアドレスを入力してください"}
	ERROR_REQUEST_PASSWORD_LATTERS  = Result{Status: 400, Message: "パスワードは8文字以上で入力してください"}
	ERROR_REQUEST_PASSWORD_STR      = Result{Status: 400, Message: "パスワードは大文字小文字数字を1つ以上含めてください"}
	ERROR_REQUEST_ACCOUNT_UNIQUE    = Result{Status: 400, Message: "既に存在するアカウントIDです"}
	ERROR_REQUEST_ACCOUNT_NOT_FOUND = Result{Status: 400, Message: "対象のアカウントが存在しません"}
	ERROR_REQUEST_FORMAT_DATE       = Result{Status: 400, Message: "YYYY-MM-DDの形式で入力してください"}
	// UnAuthorized(401)
	ERROR_AUTH_GENERAL = Result{Status: 401, Message: "認証に失敗しました"}
	ERROR_AUTH_LOGIN   = Result{Status: 401, Message: "IDまたはパスワードが違います"}
	ERROR_AUTH_INVALID = Result{Status: 401, Message: "アクセストークンが有効ではありません"}
	// Forbidden(403)
	ERROR_FORBIDDEN_GENERAL = Result{Status: 403, Message: "対象のアクセスが許可されていません"}
	// NotFound(404)
	ERROR_NOTFOUND_GENERAL = Result{Status: 404, Message: "アドレスが見つかりませんでした"}
	// InternalServerError(500)
	ERROR_INTERNAL_SERVER = Result{Status: 500, Message: "サーバー側で問題が発生しました"}
)

/* アプリケーションレスポンスキー */
var (
	// OK(200)
	SUCCESS_USER_REGISTER = Result{Status: 200, Message: "ユーザー情報を登録しました"}
	SUCCESS_USER_CHANGE   = Result{Status: 200, Message: "ユーザー情報を更新しました"}
	// BadRequest(400)
	ERROR_USER_REGIST_PRIMARY   = Result{Status: 400, Message: "既に存在するユーザーIDです"}
	ERROR_USER_CHANGE_NOT_FOUND = Result{Status: 400, Message: "変更対象のユーザーが見つかりません"}
)
