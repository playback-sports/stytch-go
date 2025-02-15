package b2c

type MagicLinksCreateParams struct {
	UserID            string     `json:"user_id,omitempty"`
	ExpirationMinutes int32      `json:"expiration_minutes,omitempty"`
	Attributes        Attributes `json:"attributes,omitempty"`
}

type MagicLinksCreateResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}

type MagicLinksAuthenticateParams struct {
	Token                  string                 `json:"token,omitempty"`
	Options                Options                `json:"options,omitempty"`
	Attributes             Attributes             `json:"attributes,omitempty"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
	CodeVerifier           string                 `json:"code_verifier,omitempty"`
}

type MagicLinksAuthenticateResponse struct {
	RequestID     string  `json:"request_id,omitempty"`
	StatusCode    int     `json:"status_code,omitempty"`
	UserID        string  `json:"user_id,omitempty"`
	MethodID      string  `json:"method_id,omitempty"`
	SessionToken  string  `json:"session_token,omitempty"`
	SessionJWT    string  `json:"session_jwt,omitempty"`
	Session       Session `json:"session,omitempty"`
	User          User    `json:"user,omitempty"`
	ResetSessions bool    `json:"reset_sessions,omitempty"`
}

// MAGIC LINK - EMAIL
type MagicLinksEmailSendParams struct {
	Email                   string     `json:"email"`
	LoginMagicLinkURL       string     `json:"login_magic_link_url,omitempty"`
	SignupMagicLinkURL      string     `json:"signup_magic_link_url,omitempty"`
	LoginExpirationMinutes  int32      `json:"login_expiration_minutes,omitempty"`
	SignupExpirationMinutes int32      `json:"signup_expiration_minutes,omitempty"`
	LoginTemplateID         string     `json:"login_template_id,omitempty"`
	SignupTemplateID        string     `json:"signup_template_id,omitempty"`
	Attributes              Attributes `json:"attributes,omitempty"`
	CodeChallenge           string     `json:"code_challenge,omitempty"`
	UserID                  string     `json:"user_id,omitempty"`
	SessionToken            string     `json:"session_token,omitempty"`
	SessionJWT              string     `json:"session_jwt,omitempty"`
	Locale                  string     `json:"locale,omitempty"`
}

type MagicLinksEmailSendResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type MagicLinksEmailLoginOrCreateParams struct {
	Email                   string     `json:"email"`
	LoginMagicLinkURL       string     `json:"login_magic_link_url,omitempty"`
	SignupMagicLinkURL      string     `json:"signup_magic_link_url,omitempty"`
	LoginExpirationMinutes  int32      `json:"login_expiration_minutes,omitempty"`
	SignupExpirationMinutes int32      `json:"signup_expiration_minutes,omitempty"`
	LoginTemplateID         string     `json:"login_template_id,omitempty"`
	SignupTemplateID        string     `json:"signup_template_id,omitempty"`
	CreateUserAsPending     bool       `json:"create_user_as_pending,omitempty"`
	Attributes              Attributes `json:"attributes,omitempty"`
	CodeChallenge           string     `json:"code_challenge,omitempty"`
	Locale                  string     `json:"locale,omitempty"`
}

type MagicLinksEmailLoginOrCreateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

type MagicLinksEmailInviteParams struct {
	Email                   string     `json:"email"`
	InviteMagicLinkURL      string     `json:"invite_magic_link_url,omitempty"`
	InviteExpirationMinutes int32      `json:"invite_expiration_minutes,omitempty"`
	InviteTemplateID        string     `json:"invite_template_id,omitempty"`
	Name                    Name       `json:"name,omitempty"`
	Attributes              Attributes `json:"attributes,omitempty"`
	Locale                  string     `json:"locale,omitempty"`
}

type MagicLinksEmailInviteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type MagicLinksEmailRevokeInviteParams struct {
	Email string `json:"email"`
}

type MagicLinksEmailRevokeInviteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}
