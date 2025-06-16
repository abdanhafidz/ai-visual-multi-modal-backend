package models

type Exception struct {
	Unauthorized               bool   `json:"unauthorized,omitempty"`
	BadRequest                 bool   `json:"bad_request,omitempty"`
	DataNotFound               bool   `json:"data_not_found,omitempty"`
	InternalServerError        bool   `json:"internal_server_error,omitempty"`
	DataDuplicate              bool   `json:"data_duplicate,omitempty"`
	QueryError                 bool   `json:"query_error,omitempty"`
	InvalidPasswordLength      bool   `json:"invalid_password_length,omitempty"`
	FailedTranscripting        bool   `json:"failed_transcripting,omitempty"`
	ReplicateConnectionRefused bool   `json:"replicated_connection_refused,omitempty"`
	AudioFileError             bool   `json:audio_file_error,omitempty`
	FailedGenerateAudio        bool   `json:audio_generation_failed,omitempty`
	Message                    string `json:"message"`
}
