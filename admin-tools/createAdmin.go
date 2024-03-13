package main

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric-ca/client"
)

func main() {
	// CA sunucusuna bağlanmak için CA'nın adresi ve TLS yapılandırmasını belirleyin
	caServerURL := "https://localhost:7054"
	caCertFile := "/path/to/ca-cert.pem"      // CA sertifikasının dosya yolu
	tlsCertFile := "/path/to/client-cert.pem" // İstemci sertifikasının dosya yolu
	tlsKeyFile := "/path/to/client-key.pem"   // İstemci özel anahtarının dosya yolu

	// CA istemcisini oluşturun
	c, err := client.NewClient(client.Config{
		URL:    caServerURL,
		MSPDir: "./msp",
		TLS: &client.TLSConfig{
			CertFiles: []string{tlsCertFile},
			KeyFile:   tlsKeyFile,
			CAFiles:   []string{caCertFile},
		},
	})

	if err != nil {
		fmt.Printf("Hata oluştu: %s\n", err)
		os.Exit(1)
	}

	// Admin kullanıcısını oluşturun
	enrollmentID := "admin"       // Kullanıcı adı
	enrollmentSecret := "adminpw" // Kullanıcı parolası

	// Admin kullanıcısını kaydettirin
	registrationRequest := &client.RegistrationRequest{
		EnrollmentID:   enrollmentID,
		Secret:         enrollmentSecret,
		Type:           "client",
		MaxEnrollments: -1, // Sınırsız kayıt izni
		Affiliation:    "", // Boş
	}
	_, err = c.Register(registrationRequest)
	if err != nil {
		fmt.Printf("Admin kullanıcısını kaydederken hata oluştu: %s\n", err)
		os.Exit(1)
	}

	// Admin kullanıcısını kaydet
	enrollmentRequest := &client.EnrollmentRequest{
		EnrollmentID: enrollmentID,
		Secret:       enrollmentSecret,
	}
	_, err = c.Enroll(enrollmentRequest)
	if err != nil {
		fmt.Printf("Admin kullanıcısını kaydederken hata oluştu: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Admin kullanıcısı başarıyla oluşturuldu ve kaydedildi.")
}
