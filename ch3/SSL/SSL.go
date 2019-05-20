package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// 生成个人使用的SSL证书及服务器私钥
func main() {
	// 算数左移,规定最大的随机数范围
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	// 获得随机序列号
	serialNumber, _ := rand.Int(rand.Reader, max)

	// 生成pki项目
	subject := pkix.Name{
		Organization:       []string{"小神童俱乐部"},
		OrganizationalUnit: []string{"学习机"},
		CommonName:         "好记性学习机",
	}

	// 产生证书模板
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	// 产生公钥
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	// DER 是 BER（basic encoding rules）下的一种编码格式，
	// 它只提供了一种编码ASN.1（abstract syntax notation. One）的语法
	// 用于加密X509文档
	// x509是一种编码过的ASN.1格式的电子文档, 记录的就是形如template所描述的信息
	derByted, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")

	// 生成pem格式文件
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derByted})
	certOut.Close()

	// 生成密钥文件
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
