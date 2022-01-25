package db

func Get(input string) (string, error) {
	return Client_redis().Get(ctx, input).Result()
}
