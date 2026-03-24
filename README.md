# go-env
go env loader

차후 확장 구조
go-env/
├── go.mod
├── LICENSE
├── README.md
├── env.go              # 메인 퍼블릭 API (패키지 진입점: 사용자는 이 파일의 함수만 호출)
├── env_test.go         # 메인 로직 테스트 (TDD를 위한 필수 요소)
│
├── parser/             # 🛠️ [확장 포인트 1] 데이터 포맷 파싱 계층
│   ├── dotenv.go       # .env 포맷 파싱 로직
│   ├── dotenv_test.go
│   # └── yaml.go       # (미래 확장) .yaml 파일 지원 시 추가
│   # └── json.go       # (미래 확장) .json 파일 지원 시 추가
│
├── source/             # 📡 [확장 포인트 2] 환경변수 공급자(Source) 계층
│   ├── file.go         # 로컬 파일 시스템에서 읽어오기
│   # └── aws_ssm.go    # (미래 확장) AWS Parameter Store에서 읽어오기
│   # └── os_env.go     # OS 시스템 환경변수와 병합(Merge) 로직
│
└── cast/               # 🔄 [확장 포인트 3] 타입 변환 유틸리티
    ├── cast.go         # GetInt(), GetBool(), GetDuration() 등
    └── cast_test.go
