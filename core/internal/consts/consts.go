package consts

const (
	DEFAULT_SERVER_NAME              = "billion-mail"
	DEFAULT_DOCKER_ENV_FILE          = "../.env"
	DEFAULT_DB_PASS_FILE             = "../DBPASS_file.pl"
	PHP_FPM_SOCK_PATH                = "../php-sock/php-fpm.sock"
	ROUNDCUBE_ROOT_PATH              = "../webmail-data"
	ROUNDCUBE_ROOT_PATH_IN_CONTAINER = "/var/www/html"
	POSTGRESQL_SOCK                  = "../postgresql-socket"
	DEFAULT_DOCKER_CLIENT_CTX_KEY    = "dockerapi"
	JWT_BLACK_LIST_KEY_PREFIX        = "JWT_BLACK_LIST:"
)
