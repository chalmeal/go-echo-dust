package response

import "log"

/*
* 要求処理に対するレスポンスを表現します。
 */
func Response(value interface{}, result Result) (int, Result) {
	return result.Status,
		Result{
			Status:  result.Status,
			Message: result.Message,
			Result:  value,
		}
}

// validateサポート
func Validate(err error) (int, Result) {
	log.Println(err)
	return 400,
		Result{
			Status:      ERROR_REQUEST_GENERAL.Status,
			Message:     ERROR_REQUEST_GENERAL.Message,
			ErrorReason: err.Error(),
			Result:      nil,
		}
}

// Success
func Success(record interface{}, response Result) Result {
	return Result{
		Status:  response.Status,
		Message: response.Message,
		Result:  record,
	}
}

// Error
func Error(err error, response Result) Result {
	log.Println(err)
	return Result{
		Status:      response.Status,
		Message:     response.Message,
		ErrorReason: err.Error(),
		Result:      nil,
	}
}
