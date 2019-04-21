package pitOrgan

// 400 Bad Request

type BadRequestError struct {
	ErrorMessage string `json:"errorMessage"`
}

func (r *BadRequestError) Error() string {
	return r.ErrorMessage
}

// 401 Unauthorized

type UnauthorizedError struct {
	ErrorMessage string `json:"errorMessage"`
}

func (r *UnauthorizedError) Error() string {
	return r.ErrorMessage
}

// 403 Forbidden

type ForbiddenError struct {
	ErrorMessage string `json:"errorMessage"`
}

func (r *ForbiddenError) Error() string {
	return r.ErrorMessage
}

// 404 Not Found

type NotFoundError struct {
	ErrorMessage string `json:"errorMessage"`
}

func (r *NotFoundError) Error() string {
	return r.ErrorMessage
}
