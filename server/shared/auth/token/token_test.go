package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const PublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu1SU1LfVLPHCozMxH2Mo
4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0/IzW7yWR7QkrmBL7jTKEn5u
+qKhbwKfBstIs+bMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyeh
kd3qqGElvW/VDL5AaWTg0nLVkjRo9z+40RQzuVaE8AkAFmxZzow3x+VJYKdjykkJ
0iT9wCS0DRTXu269V264Vf/3jvredZiKRkgwlL9xNAwxXFg0x/XFw005UWVRIkdg
cKWTjpBP2dPwVZ4WWC+9aGVd+Gyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbc
mwIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {

	pk, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PublicKey))
	if err != nil {
		t.Fatalf("cannot pares publickey:%v", err)
	}
	v := JWTTokenVerifier{
		PublicKey: pk,
	}

	cases := []struct {
		name    string
		tkn     string
		want    string
		wantErr bool
		now     time.Time
	}{
		{
			name:    "valid_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.GcHiRgOuFiiQJAMKJemV2j5Vr8uZslvOJksONETcsXxTpDqEJPHiwLsc94W3cvVpYrJO6O6c8mywxYjOWkk7iBoyEWMmbapsE8T3dDyFRq2xnV-1DZerlTNVuO4gT2fq3eNOEE-XXu0y0zlnCW7LMnOZdstHAkMD-ZQP0vKZuLJjP_AMhfd3BcsVXTMLVKjW0aG-UwkAhsathBa24NaLy2AsCIljSGNjmQ4gp9CihlHDRyUCRxBPuKDf0ym-tBSUgWk9zFugKlx-nSCYLSXgMPJ0CzgSDvmoXkC3HNM1VWOo-qd-QtInMYYuQs_RSPK8VVDj7EV7llHcjbki-OtRPA",
			want:    "1234567890",
			wantErr: false,
			now:     time.Unix(1516239022, 0),
		},
		{
			name:    "expired_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.GcHiRgOuFiiQJAMKJemV2j5Vr8uZslvOJksONETcsXxTpDqEJPHiwLsc94W3cvVpYrJO6O6c8mywxYjOWkk7iBoyEWMmbapsE8T3dDyFRq2xnV-1DZerlTNVuO4gT2fq3eNOEE-XXu0y0zlnCW7LMnOZdstHAkMD-ZQP0vKZuLJjP_AMhfd3BcsVXTMLVKjW0aG-UwkAhsathBa24NaLy2AsCIljSGNjmQ4gp9CihlHDRyUCRxBPuKDf0ym-tBSUgWk9zFugKlx-nSCYLSXgMPJ0CzgSDvmoXkC3HNM1VWOo-qd-QtInMYYuQs_RSPK8VVDj7EV7llHcjbki-OtRPA",
			wantErr: true,
			now:     time.Unix(1517239022, 0),
		},
		{
			name:    "bad_token",
			tkn:     "eyJhbGciOiJ123213234XVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.GcHiRgOuFiiQJAMKJemV2j5Vr8uZslvOJksONETcsXxTpDqEJPHiwLsc94W3cvVpYrJO6O6c8mywxYjOWkk7iBoyEWMmbapsE8T3dDyFRq2xnV-1DZerlTNVuO4gT2fq3eNOEE-XXu0y0zlnCW7LMnOZdstHAkMD-ZQP0vKZuLJjP_AMhfd3BcsVXTMLVKjW0aG-UwkAhsathBa24NaLy2AsCIljSGNjmQ4gp9CihlHDRyUCRxBPuKDf0ym-tBSUgWk9zFugKlx-nSCYLSXgMPJ0CzgSDvmoXkC3HNM1VWOo-qd-QtInMYYuQs_RSPK8VVDj7EV7llHcjbki-OtRPA",
			want:    "1234567890",
			wantErr: true,
			now:     time.Unix(1517239022, 0),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verify(c.tkn)
			if !c.wantErr && err != nil {
				if err != nil {
					t.Errorf("verification faild: %v", err)
				}
			}
			if c.wantErr && err == nil {
				if err != nil {
					t.Errorf("want error but no error")
				}
			}
			if accountID != c.want {
				t.Errorf("wrong accoid.want id:%q\ngotid:%q", c.tkn, accountID)
			}
		})
	}
	tkn := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.GcHiRgOuFiiQJAMKJemV2j5Vr8uZslvOJksONETcsXxTpDqEJPHiwLsc94W3cvVpYrJO6O6c8mywxYjOWkk7iBoyEWMmbapsE8T3dDyFRq2xnV-1DZerlTNVuO4gT2fq3eNOEE-XXu0y0zlnCW7LMnOZdstHAkMD-ZQP0vKZuLJjP_AMhfd3BcsVXTMLVKjW0aG-UwkAhsathBa24NaLy2AsCIljSGNjmQ4gp9CihlHDRyUCRxBPuKDf0ym-tBSUgWk9zFugKlx-nSCYLSXgMPJ0CzgSDvmoXkC3HNM1VWOo-qd-QtInMYYuQs_RSPK8VVDj7EV7llHcjbki-OtRPA"
	jwt.TimeFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	accountID, err := v.Verify(tkn)
	if err != nil {
		t.Errorf("verification faild: %v", err)
	}
	want := "1234567890"
	if accountID != want {
		t.Errorf("wrong accoid.want id:%q\ngotid:%q", want, accountID)
	}

}
