package sqlite

func (sq *Sqlite) SaveRefreshTokenToDB(userID int, refreshToken string) error {
	stmt := "UPDATE users SET refresh_token=? WHERE user_id=?"
	s, err := sq.db.Exec(stmt, refreshToken, userID)
	sq.log.Info(s)
	if err != nil {
		return err
	}
	return nil
}
