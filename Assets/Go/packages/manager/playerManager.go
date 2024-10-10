package manager

import (
	"encoding/binary"
	"errors"
	"log"

	pb "golangtcp/messages"

	"net"

	"google.golang.org/protobuf/proto"
)

var playerManager *PlayerManager

// Player represents a single player with some attributes
type Player struct {
	ID         int
	Name       string
	Age        int
	Conn       *net.Conn
	X          float32
	Y          float32
	Z          float32
	Rx         float32
	Ry         float32
	Rz         float32
	FinishTime int64
}

// PlayerManager manages a list of players
type PlayerManager struct {
	players map[int]*Player
	nextID  int
}

// NewPlayerManager creates a new PlayerManager
func GetPlayerManager() *PlayerManager {
	if playerManager == nil {
		playerManager = &PlayerManager{
			players: make(map[int]*Player),
			nextID:  1,
		}
	}

	return playerManager
}

// AddPlayer adds a new player to the manager
func (pm *PlayerManager) AddPlayer(name string, age int, conn *net.Conn) Player {
	player := Player{
		ID:   pm.nextID,
		Name: name,
		Age:  age,
		Conn: conn,
		X:    0,
		Y:    0,
		Z:    0,
		Rx:   0,
		Ry:   0,
		Rz:   0,
	}

	pm.players[pm.nextID] = &player
	pm.nextID++

	pm.SpawnNewPlayerInfo(player)
	pm.SendExistingPlayersToNewPlayer(player)

	return player
}

// 로그인 시 내 정보를 다른 플레이어들에게 전송해서 나를 스폰하도록 한다.
func (pm *PlayerManager) SpawnNewPlayerInfo(newPlayer Player) {
	gameMessage := &pb.GameMessage{
		Message: &pb.GameMessage_SpawnPlayer{
			SpawnPlayer: &pb.SpawnPlayer{
				PlayerId: newPlayer.Name,
				X:        0,
				Y:        0,
				Z:        0,
				Rx:       0,
				Ry:       0,
				Rz:       0,
			},
		},
	}

	// 직렬화
	response, err := proto.Marshal(gameMessage)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	// 다른 플레이어들에게 전송
	for _, player := range pm.players {
		if player.Name == newPlayer.Name {
			continue // 자신에게는 전송하지 않음
		}

		lengthBuf := make([]byte, 4)
		binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))

		// 메시지 길이 정보와 메시지 데이터를 결합하여 전송
		lengthBuf = append(lengthBuf, response...)
		(*player.Conn).Write(lengthBuf)
	}
}

func (pm *PlayerManager) SendExistingPlayersToNewPlayer(newPlayer Player) {
	for _, existingPlayer := range pm.players {
		if existingPlayer.Name == newPlayer.Name {
			// 자신은 제외
			continue
		}

		// 기존 플레이어 정보를 새로운 유저에게 전송
		gameMessage := &pb.GameMessage{
			Message: &pb.GameMessage_SpawnExistingPlayer{
				SpawnExistingPlayer: &pb.SpawnPlayer{
					PlayerId: existingPlayer.Name,
					X:        existingPlayer.X,
					Y:        existingPlayer.Y,
					Z:        existingPlayer.Z,
					Rx:       existingPlayer.Rx,
					Ry:       existingPlayer.Ry,
					Rz:       existingPlayer.Rz,
				},
			},
		}

		response, err := proto.Marshal(gameMessage)
		if err != nil {
			log.Printf("Failed to marshal response: %v", err)
			return
		}

		// 새로운 플레이어에게 기존 플레이어 정보 전송
		length := uint32(len(response))
		lengthBuf := make([]byte, 4)

		binary.LittleEndian.PutUint32(lengthBuf, length)
		if _, err := (*newPlayer.Conn).Write(append(lengthBuf, response...)); err != nil {
			log.Printf("Failed to send player data to new player: %v", err)
		}
		//binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))
		//lengthBuf = append(lengthBuf, response...)
		//(*newPlayer.Conn).Write(response)
	}
}

func (pm *PlayerManager) SendPlayerAnimation(name string, animation string, speedF float32, speedR float32) {
	// GameMessage에 PlayerAnimation 타입 추가
	gameMessage := &pb.GameMessage{
		Message: &pb.GameMessage_PlayerAnimState{
			PlayerAnimState: &pb.PlayerAnimation{
				PlayerAnimState: animation,
				PlayerId:        name,
				SpeedForward:    speedF,
				SpeedRight:      speedR,
			},
		},
	}

	// 직렬화
	response, err := proto.Marshal(gameMessage)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	// 다른 플레이어들에게 전송
	for _, player := range pm.players {
		if player.Name == name {
			continue // 자신에게는 전송하지 않음
		}

		lengthBuf := make([]byte, 4)
		binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))

		// 메시지 길이 정보와 메시지 데이터를 결합하여 전송
		lengthBuf = append(lengthBuf, response...)
		(*player.Conn).Write(lengthBuf)
	}
}

func (pm *PlayerManager) MovePlayer(name string, x float32, y float32, z float32, rx float32, ry float32, rz float32) {
	player, exists := pm.FindPlayerByName(name)
	if !exists {
		log.Printf("Player not found: %s", name)
		return
	}

	// 플레이어의 위치 업데이트
	player.X = x
	player.Y = y
	player.Z = z
	player.Rx = rx
	player.Ry = ry
	player.Rz = rz

	gameMessage := &pb.GameMessage{
		Message: &pb.GameMessage_PlayerPosition{
			PlayerPosition: &pb.PlayerPosition{
				PlayerId: name,
				X:        x,
				Y:        y,
				Z:        z,
				Rx:       rx,
				Ry:       ry,
				Rz:       rz,
			},
		},
	}

	response, err := proto.Marshal(gameMessage)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	for _, player := range pm.players {
		if player.Name == name {
			continue
		}

		lengthBuf := make([]byte, 4)
		binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))
		lengthBuf = append(lengthBuf, response...)
		(*player.Conn).Write(lengthBuf)
	}
}

// GetPlayer retrieves a player by ID
func (pm *PlayerManager) GetPlayer(id int) (*Player, error) {
	player, exists := pm.players[id]
	if !exists {
		return nil, errors.New("player not found")
	}
	return player, nil
}

// RemovePlayer removes a player by ID
func (pm *PlayerManager) RemovePlayer(id string) error {
	player, exists := pm.FindPlayerByName(id)
	if !exists {
		log.Printf("Player not found: %s", id)
		return errors.New("Player not found")
	}
	delete(pm.players, player.ID)

	logoutPacket := &pb.GameMessage{
		Message: &pb.GameMessage_Logout{
			Logout: &pb.LogoutMessage{
				PlayerId: id,
			},
		},
	}

	response, err := proto.Marshal(logoutPacket)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return errors.New("Fail to send logout packet")
	}

	for _, player := range pm.players {
		if player.Name == id {
			continue // 자신에게는 전송하지 않음
		}

		lengthBuf := make([]byte, 4)
		binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))

		// 메시지 길이 정보와 메시지 데이터를 결합하여 전송
		lengthBuf = append(lengthBuf, response...)
		(*player.Conn).Write(lengthBuf)
	}

	// // 이 코드를 들어온 유저를 제외한 플레이어들에게 스폰시켜달라고 한다.
	// for _, p := range pm.players {
	// 	(*p.Conn).Write(response)
	// }

	return nil
}

// ListPlayers returns all players in the manager
func (pm *PlayerManager) ListPlayers() []*Player {
	playerList := []*Player{}
	for _, player := range pm.players {
		playerList = append(playerList, player)
	}
	return playerList
}

func (pm *PlayerManager) FindPlayerByName(name string) (*Player, bool) {
	for _, player := range pm.players {
		if player.Name == name {
			return player, true // 포인터를 반환합니다.
		}
	}
	return nil, false // 찾지 못한 경우 nil과 false를 반환합니다.
}

func (pm *PlayerManager) BroadcastMessage(message *pb.GameMessage) {
	response, err := proto.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	for _, player := range pm.players {
		lengthBuf := make([]byte, 4)
		binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))

		// 메시지 길이 정보와 메시지 데이터를 결합하여 전송
		lengthBuf = append(lengthBuf, response...)
		(*player.Conn).Write(lengthBuf)
	}
}

func (pm *PlayerManager) PlayerFinishedRace(playerId string, finishTime int64) {
	player, exists := pm.FindPlayerByName(playerId)
	if !exists {
		log.Printf("Player not found: %s", playerId)
		return
	}

	player.FinishTime = finishTime

	// 모든 플레이어에게 완주 정보 브로드캐스트
	finishMessage := &pb.GameMessage{
		Message: &pb.GameMessage_RaceFinish{
			RaceFinish: &pb.RaceFinishMessage{
				PlayerId:   playerId,
				FinishTime: finishTime,
			},
		},
	}

	pm.BroadcastMessage(finishMessage)
}
