package postgres

func (c *Connection) FlushDB() {
	stmt, _ := c.db.Prepare("DELETE FROM users")
	_, _ = stmt.Exec()

	stmt, _ = c.db.Prepare("DELETE FROM auth")
	_, _ = stmt.Exec()
	return
}
