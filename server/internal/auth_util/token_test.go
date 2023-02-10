package auth

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const PUBLIC_KEY = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAro6z1vz7d4HL6v9u+gLm
oOKLFwcQS5k2uKRpGvLvxcflOp3nYUbnuURnakyzjxVTa+V2887039Jjp5SQzaxy
U02sym4ZWPcwvZsGDCqMN2EC9oeaXdqKYNNETGrZwkvOV2lYD/GVgBYy3sVW7/pz
dUA772+hmY1nHl5EOFkcK/Ds56bDh9jKTJci99ebTqbuYYKKYgphEHE6GuA/gpKr
i/qIE93/Zbpc/KNOIl9ncquyTdktTHodVZKMTzfvMfsmNlCW0w9DLY7Xi2qYFALa
e0/uWCrvmds7k2NHNDPSPzEldE+YyINnj2KJ9L3BVoHImbrOA0isF1dOG5yN8EZl
eQIDAQAB
-----END PUBLIC KEY-----`

func TestJWTTokenVerifier_Verify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PUBLIC_KEY))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string
		token   string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name:    "valid_token",
			token:   `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI2M2U1MDg0YmI2M2YzZmZhOTk4OTIwNDgiLCJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoic2ZjYXIvYXV0aCJ9.OWKRPE73ZFhsghdUzn5dZZHK_aIrMfHgKtzt1b45tJbUqD1_bNkSRUdAl86x8be3UHso4IYzR9wTXySE5-TJVwzrzT-7MdXdVIV4qq8T1wRDnWg-xV6dff4acEeDIN3gPwQ4Ihx5sdIoMQ6r8hTSQ4wYtUjOQB8f4dgf9YwMpG-OAenXkmA3Fe8zEuK1YyILucUlSLI9MNL165xopXQTax_LYZM8mNd3YzWIaJ_fKg1c5kgt3Qjjrf5EoBOOKN3Y7h2_vZ9D6nhltha0wyeCdaKJQmecygV7N87jkCxeqOhs3N-XWLSMwPWcuTFFrl2y_itQWbAgmdN8eCOuCAToEg`,
			now:     time.Unix(1516239032, 0),
			want:    "63e5084bb63f3ffa99892048",
			wantErr: false,
		},
		{
			name:    "token_expired",
			token:   `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI2M2U1MDg0YmI2M2YzZmZhOTk4OTIwNDgiLCJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoic2ZjYXIvYXV0aCJ9.OWKRPE73ZFhsghdUzn5dZZHK_aIrMfHgKtzt1b45tJbUqD1_bNkSRUdAl86x8be3UHso4IYzR9wTXySE5-TJVwzrzT-7MdXdVIV4qq8T1wRDnWg-xV6dff4acEeDIN3gPwQ4Ihx5sdIoMQ6r8hTSQ4wYtUjOQB8f4dgf9YwMpG-OAenXkmA3Fe8zEuK1YyILucUlSLI9MNL165xopXQTax_LYZM8mNd3YzWIaJ_fKg1c5kgt3Qjjrf5EoBOOKN3Y7h2_vZ9D6nhltha0wyeCdaKJQmecygV7N87jkCxeqOhs3N-XWLSMwPWcuTFFrl2y_itQWbAgmdN8eCOuCAToEg`,
			now:     time.Unix(2516239022, 0),
			want:    "",
			wantErr: true,
		},
		{
			name:    "bad_token",
			token:   "bad_token",
			now:     time.Unix(1516239032, 0),
			want:    "",
			wantErr: true,
		},
		{
			name:    "wrong_signature",
			token:   `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI2M2U1MDg0YmI2M2YzZmZhOTk4OTIwNDkiLCJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoic2ZjYXIvYXV0aCJ9.OWKRPE73ZFhsghdUzn5dZZHK_aIrMfHgKtzt1b45tJbUqD1_bNkSRUdAl86x8be3UHso4IYzR9wTXySE5-TJVwzrzT-7MdXdVIV4qq8T1wRDnWg-xV6dff4acEeDIN3gPwQ4Ihx5sdIoMQ6r8hTSQ4wYtUjOQB8f4dgf9YwMpG-OAenXkmA3Fe8zEuK1YyILucUlSLI9MNL165xopXQTax_LYZM8mNd3YzWIaJ_fKg1c5kgt3Qjjrf5EoBOOKN3Y7h2_vZ9D6nhltha0wyeCdaKJQmecygV7N87jkCxeqOhs3N-XWLSMwPWcuTFFrl2y_itQWbAgmdN8eCOuCAToEg`,
			now:     time.Unix(1516239032, 0),
			want:    "",
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}

			accountID, err := v.Verify(c.token)
			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}
			if c.wantErr && err == nil {
				t.Errorf("want error; got no error")
			}

			if accountID != c.want {
				t.Errorf("wrong account id.\n want: %q,\n got: %q", c.want, accountID)
			}
		})
	}
}
