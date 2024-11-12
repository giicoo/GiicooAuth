package services

func (s *Services) GenerateJWT(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if s.hashTools.CheckPasswordHash(password, user.HashPassword) {
		jwtToken, err := s.jwtTools.NewJWT(user.UserId, user.Email)
		if err != nil {
			return "", err
		}
		return jwtToken, nil
	}
	return "", err
}

func (s *Services) CheckJWT(jwtToken string) (int, string, error) {
	return s.jwtTools.ParseJWT(jwtToken)
}
