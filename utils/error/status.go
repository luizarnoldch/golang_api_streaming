package error

// HTTP status codes were copied from net/http with the following updates:
// - Rename StatusNonAuthoritativeInfo to StatusNonAuthoritativeInformation
// - Add StatusSwitchProxy (306)
// NOTE: Keep this list in sync with statusMessage
const (
	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                          = 200 // RFC 9110, 15.3.1
	StatusCreated                     = 201 // RFC 9110, 15.3.2
	StatusAccepted                    = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInformation = 203 // RFC 9110, 15.3.4
	StatusNoContent                   = 204 // RFC 9110, 15.3.5
	StatusResetContent                = 205 // RFC 9110, 15.3.6
	StatusPartialContent              = 206 // RFC 9110, 15.3.7
	StatusMultiStatus                 = 207 // RFC 4918, 11.1
	StatusAlreadyReported             = 208 // RFC 5842, 7.1
	StatusIMUsed                      = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently  = 301 // RFC 9110, 15.4.2
	StatusFound             = 302 // RFC 9110, 15.4.3
	StatusSeeOther          = 303 // RFC 9110, 15.4.4
	StatusNotModified       = 304 // RFC 9110, 15.4.5
	StatusUseProxy          = 305 // RFC 9110, 15.4.6
	StatusSwitchProxy       = 306 // RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

// Errors
// var (
// 	ErrBadRequest                   = NewError(StatusBadRequest)                   // 400
// 	ErrUnauthorized                 = NewError(StatusUnauthorized)                 // 401
// 	ErrPaymentRequired              = NewError(StatusPaymentRequired)              // 402
// 	ErrForbidden                    = NewError(StatusForbidden)                    // 403
// 	ErrNotFound                     = NewError(StatusNotFound)                     // 404
// 	ErrMethodNotAllowed             = NewError(StatusMethodNotAllowed)             // 405
// 	ErrNotAcceptable                = NewError(StatusNotAcceptable)                // 406
// 	ErrProxyAuthRequired            = NewError(StatusProxyAuthRequired)            // 407
// 	ErrRequestTimeout               = NewError(StatusRequestTimeout)               // 408
// 	ErrConflict                     = NewError(StatusConflict)                     // 409
// 	ErrGone                         = NewError(StatusGone)                         // 410
// 	ErrLengthRequired               = NewError(StatusLengthRequired)               // 411
// 	ErrPreconditionFailed           = NewError(StatusPreconditionFailed)           // 412
// 	ErrRequestEntityTooLarge        = NewError(StatusRequestEntityTooLarge)        // 413
// 	ErrRequestURITooLong            = NewError(StatusRequestURITooLong)            // 414
// 	ErrUnsupportedMediaType         = NewError(StatusUnsupportedMediaType)         // 415
// 	ErrRequestedRangeNotSatisfiable = NewError(StatusRequestedRangeNotSatisfiable) // 416
// 	ErrExpectationFailed            = NewError(StatusExpectationFailed)            // 417
// 	ErrTeapot                       = NewError(StatusTeapot)                       // 418
// 	ErrMisdirectedRequest           = NewError(StatusMisdirectedRequest)           // 421
// 	ErrUnprocessableEntity          = NewError(StatusUnprocessableEntity)          // 422
// 	ErrLocked                       = NewError(StatusLocked)                       // 423
// 	ErrFailedDependency             = NewError(StatusFailedDependency)             // 424
// 	ErrTooEarly                     = NewError(StatusTooEarly)                     // 425
// 	ErrUpgradeRequired              = NewError(StatusUpgradeRequired)              // 426
// 	ErrPreconditionRequired         = NewError(StatusPreconditionRequired)         // 428
// 	ErrTooManyRequests              = NewError(StatusTooManyRequests)              // 429
// 	ErrRequestHeaderFieldsTooLarge  = NewError(StatusRequestHeaderFieldsTooLarge)  // 431
// 	ErrUnavailableForLegalReasons   = NewError(StatusUnavailableForLegalReasons)   // 451

// 	ErrInternalServerError           = NewError(StatusInternalServerError)           // 500
// 	ErrNotImplemented                = NewError(StatusNotImplemented)                // 501
// 	ErrBadGateway                    = NewError(StatusBadGateway)                    // 502
// 	ErrServiceUnavailable            = NewError(StatusServiceUnavailable)            // 503
// 	ErrGatewayTimeout                = NewError(StatusGatewayTimeout)                // 504
// 	ErrHTTPVersionNotSupported       = NewError(StatusHTTPVersionNotSupported)       // 505
// 	ErrVariantAlsoNegotiates         = NewError(StatusVariantAlsoNegotiates)         // 506
// 	ErrInsufficientStorage           = NewError(StatusInsufficientStorage)           // 507
// 	ErrLoopDetected                  = NewError(StatusLoopDetected)                  // 508
// 	ErrNotExtended                   = NewError(StatusNotExtended)                   // 510
// 	ErrNetworkAuthenticationRequired = NewError(StatusNetworkAuthenticationRequired) // 511
// )