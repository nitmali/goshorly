package db

func Get(input string) (string, error) {
	return Client.Get(ctx, input).Result()
}
