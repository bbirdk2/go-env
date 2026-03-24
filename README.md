# 📦 go-env

`go-env`는 외부 의존성(Zero-Dependency) 없이 Go 언어 내장 패키지만을 활용하여 설계된 가볍고 확장 가능한 환경변수 로더(Environment Variable Loader)입니다.

B2B SaaS 및 마이크로서비스 아키텍처 환경에서 빠르고 안전하게 설정값을 주입하기 위해 만들어졌습니다.

## 🚀 Installation

```bash
go get [github.com/bbirdk2/go-env](https://github.com/bbirdk2/go-env)


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
