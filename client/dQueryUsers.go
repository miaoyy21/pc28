package client

import (
	"database/sql"
)

type User struct {
	UserName string
	IsMaster bool
	Host     string
	Sigma    float64

	Cookie    string
	UserAgent string
	Unix      string
	KeyCode   string
	DeviceId  string
	UserId    string
	Token     string

	Gold int64
}

func dQueryUsers(db *sql.DB) ([]*User, error) {
	query := `
		SELECT user_name, is_master, host, sigma, cookie, user_agent, unix, key_code, device_id, user_id, token, gold
		FROM user
		WHERE is_valid = 1
		ORDER BY gold DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var userName, host, cookie, userAgent, unix, keyCode, deviceId, userId, token string
		var isMaster bool
		var sigma float64
		var gold int64
		if err := rows.Scan(&userName, &isMaster, &host, &sigma, &cookie, &userAgent, &unix, &keyCode, &deviceId, &userId, &token, &gold); err != nil {
			return nil, err
		}

		user := &User{
			UserName: userName,
			IsMaster: isMaster,
			Host:     host,
			Sigma:    sigma,

			Cookie:    cookie,
			UserAgent: userAgent,
			Unix:      unix,
			KeyCode:   keyCode,
			DeviceId:  deviceId,
			UserId:    userId,
			Token:     token,

			Gold: gold,
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
