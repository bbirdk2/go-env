# 📦 go-env

`go-env`는 외부 의존성(Zero-Dependency) 없이 Go 언어 내장 패키지만을 활용하여 설계된 가볍고 확장 가능한 환경변수 로더(Environment Variable Loader)입니다.

B2B SaaS 및 마이크로서비스 아키텍처 환경에서 빠르고 안전하게 설정값을 주입하기 위해 만들어졌습니다.

## 🚀 Installation

```bash
go get github.com/bbirdk2/go-env(https://github.com/bbirdk2/go-env)
```

## 📂 Directory Architecture

```bash
go-env/
├── go.mod
├── LICENSE
├── README.md
├── env.go              # 🚪 메인 퍼블릭 API (패키지 진입점: Facade 역할)
├── env_test.go         # ✅ 메인 로직 테스트
│
├── parser/             # 🛠️ [확장 포인트 1] 데이터 포맷 파싱 계층
│   ├── dotenv.go       # .env 포맷 파싱 로직
│   └── dotenv_test.go
│
├── source/             # 📡 [확장 포인트 2] 환경변수 공급자(Source) 계층
│   └── file.go         # 로컬 파일 시스템 리더 (향후 AWS SSM 등으로 확장 가능)
│
└── cast/               # 🔄 [확장 포인트 3] 타입 변환 유틸리티
    ├── cast.go         # GetInt, GetBool 등 형변환 및 기본값(Fallback) 지원
    └── cast_test.go
```

## 💡 구조 설계 철학 (Design Principles)
```bash
Facade Pattern (env.go): 사용자는 내부의 복잡한 패키지 구조를 알 필요 없이 최상단의 env.go를 통해 직관적으로 기능을 호출합니다.

Open-Closed Principle (parser/): 기존 코드를 수정하지 않고도 새로운 설정 파일 포맷(예: config.yaml)을 지원할 수 있도록 파싱 로직을 분리했습니다.

Type Casting (cast/): 실무에서 번거로운 string -> int/bool 변환 로직을 내재화하여 비즈니스 로직의 피로도를 낮춥니다.
```

## 💻 Quick Start
```bash
package main

import (
	"fmt"
	"github.com/bbirdk2/go-env(https://github.com/bbirdk2/go-env)"
)

func main() {
	// 1. .env 파일 로드
	err := env.Load(".env")
	if err != nil {
		fmt.Println("Warning: .env 파일을 찾을 수 없습니다.")
	}

	// 2. 환경변수 읽기 (기본값 설정 가능)
	apiKey := env.Get("API_KEY", "default-key-123")
	
	// 3. 타입 변환하여 읽기 (향후 cast 패키지 구현 시)
	// port := env.GetInt("PORT", 8080)
	// isDebug := env.GetBool("DEBUG_MODE", false)

	fmt.Printf("Loaded API Key: %s\n", apiKey)
}
```

## 📄 License
This project is licensed under the MIT License.
