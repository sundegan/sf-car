package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const PRIVATE_KEY = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC7VJTUt9Us8cKj
MzEfYyjiWA4R4/M2bS1GB4t7NXp98C3SC6dVMvDuictGeurT8jNbvJZHtCSuYEvu
NMoSfm76oqFvAp8Gy0iz5sxjZmSnXyCdPEovGhLa0VzMaQ8s+CLOyS56YyCFGeJZ
qgtzJ6GR3eqoYSW9b9UMvkBpZODSctWSNGj3P7jRFDO5VoTwCQAWbFnOjDfH5Ulg
p2PKSQnSJP3AJLQNFNe7br1XbrhV//eO+t51mIpGSDCUv3E0DDFcWDTH9cXDTTlR
ZVEiR2BwpZOOkE/Z0/BVnhZYL71oZV34bKfWjQIt6V/isSMahdsAASACp4ZTGtwi
VuNd9tybAgMBAAECggEBAKTmjaS6tkK8BlPXClTQ2vpz/N6uxDeS35mXpqasqskV
laAidgg/sWqpjXDbXr93otIMLlWsM+X0CqMDgSXKejLS2jx4GDjI1ZTXg++0AMJ8
sJ74pWzVDOfmCEQ/7wXs3+cbnXhKriO8Z036q92Qc1+N87SI38nkGa0ABH9CN83H
mQqt4fB7UdHzuIRe/me2PGhIq5ZBzj6h3BpoPGzEP+x3l9YmK8t/1cN0pqI+dQwY
dgfGjackLu/2qH80MCF7IyQaseZUOJyKrCLtSD/Iixv/hzDEUPfOCjFDgTpzf3cw
ta8+oE4wHCo1iI1/4TlPkwmXx4qSXtmw4aQPz7IDQvECgYEA8KNThCO2gsC2I9PQ
DM/8Cw0O983WCDY+oi+7JPiNAJwv5DYBqEZB1QYdj06YD16XlC/HAZMsMku1na2T
N0driwenQQWzoev3g2S7gRDoS/FCJSI3jJ+kjgtaA7Qmzlgk1TxODN+G1H91HW7t
0l7VnL27IWyYo2qRRK3jzxqUiPUCgYEAx0oQs2reBQGMVZnApD1jeq7n4MvNLcPv
t8b/eU9iUv6Y4Mj0Suo/AU8lYZXm8ubbqAlwz2VSVunD2tOplHyMUrtCtObAfVDU
AhCndKaA9gApgfb3xw1IKbuQ1u4IF1FJl3VtumfQn//LiH1B3rXhcdyo3/vIttEk
48RakUKClU8CgYEAzV7W3COOlDDcQd935DdtKBFRAPRPAlspQUnzMi5eSHMD/ISL
DY5IiQHbIH83D4bvXq0X7qQoSBSNP7Dvv3HYuqMhf0DaegrlBuJllFVVq9qPVRnK
xt1Il2HgxOBvbhOT+9in1BzA+YJ99UzC85O0Qz06A+CmtHEy4aZ2kj5hHjECgYEA
mNS4+A8Fkss8Js1RieK2LniBxMgmYml3pfVLKGnzmng7H2+cwPLhPIzIuwytXywh
2bzbsYEfYx3EoEVgMEpPhoarQnYPukrJO4gwE2o5Te6T5mJSZGlQJQj9q4ZB2Dfz
et6INsK0oG8XVGXSpQvQh3RUYekCZQkBBFcpqWpbIEsCgYAnM3DQf3FJoSnXaMhr
VBIovic5l0xFkEHskAjFTevO86Fsz1C2aSeRKSqGFoOQ0tmJzBEs1R6KqnHInicD
TQrKhArgLXX4v3CddjfTRJkFWDbE/CkvKZNOrcf1nhaGCPspRJj2KUkj1Fhl9Cnc
dn/RsYEONbwQSjIfMPkvxF+8HQ==
-----END PRIVATE KEY-----`

func TestJWTTokenGen_GenerateToken(t *testing.T) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(PRIVATE_KEY))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := &JWTTokenGen{
		PrivateKey: privateKey,
		Issuer:     "sfcar/auth",
		IssuedAt: func() time.Time {
			return time.Unix(1516239022, 0)
		},
	}
	token, err := g.GenerateToken("63e26d0625d9b723e3f81900", time.Minute)
	if err != nil {
		t.Errorf("cannot generate token: %v", err)
	}
	want := `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI2M2UyNmQwNjI1ZDliNzIzZTNmODE5MDAiLCJleHAiOjE1MTYyMzkwODIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoic2ZjYXIvYXV0aCJ9.p90FqvbiiXpBPwa5mTOR1vJo-4mf19TU9N09irxMl0pHtR-8f5s6xGDBmoHCbiCzaheavf0PSfBnsbM54JBEx6fhG4Squh_jGQ1Tvdn5NBtdCu1gmklKluqnSvlHi76ApdEbz1VAnjFJlCF-182yndLh0yaXrofenwI0eLvz5UYsahUCt8F9MHsCDmeCIggY3VbYVo24UEsmtlXDkWUUIrpvYgWlil_h-kT57HyovNi_r4Oj3_u715vFcIw-3b-pqnFOfaVcOF6zKpxRgBA5_eR_x40gipUa11cqcLGCAOtucTG-axkX8a21OBXQBFNRnX3k4EBYnWzbSB7rdpF78Q`
	if token != want {
		t.Errorf("wrong token generated,\n want: %q,\n got: %q\n", want, token)
	}
}
