package env // 패키지명

import (
	"bufio"
	"os"
	"strings"
)

// Load 함수: .env 파일을 읽어서 시스템 환경변수에 등록합니다.
func Load(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
			os.Setenv(key, value)
		}
	}
	return scanner.Err()
}

// Get 함수: 환경변수를 가져오되, 없을 경우 기본값을 반환하는 편의성 함수
func Get(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}