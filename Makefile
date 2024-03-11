SRC_DIR = ./server/src/
COMMON_DIR = ./common/
SRC = utils.go hitbox.go player.go object.go game.go main.go
COMMON = packet.go

ISHANK_SRC = utils.go main.go
ELLIOT_SRC = packet.go utils.go hitbox.go player.go object.go game.go main.go
# ELLIOT_COMMON = packet.go

default:
	tsc
#	go run $(addprefix $(COMMON_DIR),$(COMMON)) $(addprefix $(SRC_DIR),$(SRC))
	go run $(addprefix $(SRC_DIR),$(SRC))

ishank:
	tsc
	go run $(addprefix $(SRC_DIR),$(ISHANK_SRC))

elliot:
	tsc
	go run $(addprefix $(SRC_DIR),$(SRC))