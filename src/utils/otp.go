package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	configure "github.com/mrzack99s/cloud-coco/src/configures"
)

type OTPInstance struct {
	Identity  string `json:"identity"`
	OTP       string `json:"otp"`
	TimeoutAt string `json:"timeout_at"`
}

func CreateOTP(identity string) OTPInstance {
	now := time.Now()
	timeOut := now.Add(time.Minute * 3)

	newOtp := OTPInstance{
		Identity:  identity,
		OTP:       otpGenerator(),
		TimeoutAt: timeOut.Format(time.RFC3339),
	}

	marshalOtp, _ := json.Marshal(newOtp)

	configure.CacheInstance().Set(fmt.Sprintf("grpa_otp:%s", newOtp.Identity), string(marshalOtp), time.Minute*3)

	return newOtp
}

func VerifyOTP(identity, otp string) (bool, error) {
	key := fmt.Sprintf("grpa_otp:%s", identity)
	if RedisFindExistingKey(key) {
		now := time.Now()

		otpInstanceStr, _ := configure.CacheInstance().Get(key).Result()

		RedisDeleteWithKey(key)

		otpInstance := OTPInstance{}
		json.Unmarshal([]byte(otpInstanceStr), &otpInstance)

		timeoutAt, err := time.Parse(time.RFC3339, otpInstance.TimeoutAt)
		if err != nil {
			return false, err
		}

		if now.Sub(timeoutAt).Minutes() > 3 {
			return false, errors.New("otp timeout")
		} else {
			if otpInstance.OTP == otp {
				return true, nil
			}
		}
		return false, nil

	} else {
		return false, errors.New("not found identity")
	}

}

func otpGenerator() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("0123456789")
	var b strings.Builder
	for i := 0; i < 6; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
