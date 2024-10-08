using System;
using System.Collections;
using System.Collections.Generic;
using Game;
using Unity.Collections;
using UnityEngine;
using UnityEngine.Serialization;

public class PlayerController : MonoBehaviour
{
    public static PlayerController Instance { get; private set; }

    private PlayerTCP myPlayerTcp;
    private Dictionary<string, OtherPlayerTCP> _otherPlayers = new();

    public PlayerTCP myPlayerTcpTemplate;
    public OtherPlayerTCP otherPlayerTcpTemplate;
    
    private void Awake()
    {
        if (Instance == null)
        {
            Instance = this;
            DontDestroyOnLoad(this.gameObject);
        }
        else
        {
            Destroy(gameObject);
        }
    }

    void Start()
    {
        GameObject SpawnPlayer = GameObject.Instantiate(myPlayerTcpTemplate.gameObject, Vector3.zero, Quaternion.identity);
        myPlayerTcp = SpawnPlayer.GetComponent<PlayerTCP>();
    }

    /*public void OnRecevieChatMsg(ChatMessage chatmsg) //유저 간의 채팅 기능
    {
        UIManager.Instance.OnRecevieChatMsg(chatmsg);
    }*/
    
    public void OnOtherPlayerPositionUpdate(PlayerPosition playerPosition)
    {
        if (_otherPlayers.TryGetValue(
                playerPosition.PlayerId, out OtherPlayerTCP otherPlayer))
        {
            otherPlayer.destination = new Vector3(playerPosition.X, playerPosition.Y, playerPosition.Z);
            otherPlayer.OtherRot = new Vector3(playerPosition.Rx, playerPosition.Ry, playerPosition.Rz);
            return;
        }
    }

    public void SpawnOtherPlayer(SpawnPlayer serverPlayer)
    {
        GameObject SpawnPlayer = GameObject.Instantiate(otherPlayerTcpTemplate.gameObject, Vector3.zero, Quaternion.identity);
        OtherPlayerTCP otherPlayerTcp = SpawnPlayer.GetComponent<OtherPlayerTCP>();

        otherPlayerTcp.destination = new Vector3(serverPlayer.X, serverPlayer.Y, serverPlayer.Z);
        otherPlayerTcp.OtherRot = new Vector3(serverPlayer.Rx, serverPlayer.Ry, serverPlayer.Rz);
        
        _otherPlayers.TryAdd(serverPlayer.PlayerId, otherPlayerTcp);
    }


    public void OnOtherPlayerAnimationStateUpdate(PlayerAnimation playerAnimation)
    {
        if (_otherPlayers.TryGetValue(
                playerAnimation.PlayerId, out OtherPlayerTCP otherPlayer))
        {
            otherPlayer.AnimTrigger(playerAnimation);
        }
    }

    public void DespawnOtherPlayer(string playerId)
    {
        if (_otherPlayers.TryGetValue(playerId, out OtherPlayerTCP otherPlayer))
        {
            Destroy(otherPlayer.gameObject);
            _otherPlayers.Remove(playerId);
        }
    }
}