package expr

// ParseAndEval will return the evaluation of the string
func ParseAndEval(s string, env Env) (string, error) {
	expr, err := Parse(s)
	if err != nil {
		return "", err
	}
	return expr.Eval(env), nil
}
