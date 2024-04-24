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

RUN pip3 install -r requirements.txt

# Runner stage
FROM scratch AS runner
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/inu-backyard /

EXPOSE 3001
CMD ["/inu-backyard"]
