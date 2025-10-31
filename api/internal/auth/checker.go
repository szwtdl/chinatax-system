package auth

type PermissionChecker interface {
	CheckPermission(userID uint, path, method string) bool
}
