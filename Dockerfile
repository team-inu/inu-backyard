#          _                   _            _             _           _            _
#         / /\                /\ \         /\ \     _    /\ \        /\ \         /\_\
#        / /  \              /  \ \       /  \ \   /\_\ /  \ \      /  \ \       / / /  _
#       / / /\ \            / /\ \ \     / /\ \ \_/ / // /\ \ \    / /\ \ \     / / /  /\_\
#      / / /\ \ \          / / /\ \ \   / / /\ \___/ // / /\ \_\  / / /\ \ \   / / /__/ / /
#     / / /  \ \ \        / / /  \ \_\ / / /  \/____// / /_/ / / / / /  \ \_\ / /\_____/ /
#    / / /___/ /\ \      / / /   / / // / /    / / // / /__\/ / / / /   / / // /\_______/
#   / / /_____/ /\ \    / / /   / / // / /    / / // / /_____/ / / /   / / // / /\ \ \
#  / /_________/\ \ \  / / /___/ / // / /    / / // / /\ \ \  / / /___/ / // / /  \ \ \
# / / /_       __\ \_\/ / /____\/ // / /    / / // / /  \ \ \/ / /____\/ // / /    \ \ \
# \_\___\     /____/_/\/_________/ \/_/     \/_/ \/_/    \_\/\/_________/ \/_/      \_\_\
# Developed by AonrokZa1392
# ติดต่อแก้สคริปได้ที่เฟส AonrokZa1392 ไม่เข้ารหัสไฟล์ support ตลอดการใช้งาน


# Build stage
FROM golang:1.22 AS builder
WORKDIR /app

COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
    -ldflags="-w -s" \
    -o ./inu-backyard ./cmd/http_server/main.go

# Runner stage
FROM alpine:3.19 AS runner
WORKDIR /app

RUN apk add python3 py3-pip
RUN apk add py3-scipy py3-scikit-learn

COPY --from=builder /app/requirements.txt .
RUN pip3 install $(grep -vE "scikit-learn|scipy" /app/requirements.txt) --break-system-packages

COPY --from=builder /app/inu-backyard /

EXPOSE 3001
CMD ["/inu-backyard"]
