using System;
using System.Collections;
using System.Collections.Generic;
using Game;
using UnityEditor;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;
using Random = UnityEngine.Random;



public class SceneChanger : MonoBehaviour
{
    public static SceneChanger Instance { get; private set; }

    /*public Button StartButton;
    public Button CustomizeButton;*/
    public List<string> RaceList = new(); //일반 레이스
    public List<string> TeamList; //팀전
    public List<string> finalList; //결승전
    public bool isRacing = false; //플레이하는 중이 아니라면 플레이어의 입력을 못 받도록 하는 변수
    public int matchingSeed = 0; //게임마다 플레이어들에게 같은 랜덤 시드를 주기 위해 서버로부터 받아서 사용

    public List<string> raceToPlay; //여기서 맵을 랜덤하게 뽑아서 경기
    public int raceToPlayIdx = 0;
    private List<string> teamToPlay; 
    private List<string> finalToPlay;
    private Dictionary<Button, Action> buttonActions = new ();
    private Tuple<int, int> matchingStatus; 
    
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

    private void OnEnable()
    {
        //플레이 할 맵 초기화
        //raceToPlay = RaceList;
    }

    private void OnDisable()
    {
        //버튼 이벤트 해제
        ClearAllButtons();
    }

    private void Update()
    {
        if (Input.GetKeyDown(KeyCode.Escape) && GetCurrentScene() != "Main" && !isRacing)
        {
            if (GetCurrentScene() == "Matching")
            {
                TcpProtobufClient.Instance.SendMatchingRequest(TCPManager.playerId, false);       
            }
            SceneManager.LoadScene("Main");
        }
    }

    public void MatchingGame()
    {
        SceneManager.LoadScene("Matching");
        TcpProtobufClient.Instance.SendMatchingRequest(TCPManager.playerId,true);
    }

    public string GetCurrentScene()
    {
        return SceneManager.GetActiveScene().name;
    }
    
    //버튼 이벤트 등록
    public void RegisterButton(Button button, Action onClickAction)
    {
        if (buttonActions.ContainsKey(button)) return;

        buttonActions.Add(button, onClickAction);
        button.onClick.AddListener(() => onClickAction?.Invoke());
    }

    // 모든 버튼의 리스너 제거
    public void ClearAllButtons()
    {
        foreach (var kvp in buttonActions)
        {
            kvp.Key.onClick.RemoveAllListeners();
        }
        buttonActions.Clear();
    }

    public void SetRaceMaps(MatchingResponse mr)
    {
        Debug.Log($"Setting race maps. Count: {mr.MapName.Count}");
        raceToPlay = new List<string>(mr.MapName);
        matchingSeed = mr.MatchingSeed;
        raceToPlayIdx = 0; // 인덱스 초기화 추가
        PlayRace();
    }
    
    //Start Button Event
    public void PlayRace()
    {
        Debug.Log($"PlayRace called. Current index: {raceToPlayIdx}, Total maps: {raceToPlay?.Count}");
    
        if (raceToPlay == null || raceToPlay.Count == 0)
        {
            Debug.LogError("No race maps available!");
            return;
        }

        if (raceToPlayIdx >= raceToPlay.Count)
        {
            Debug.LogError($"Race index out of range: {raceToPlayIdx} / {raceToPlay.Count}");
            return;
        }

        string mapToLoad = raceToPlay[raceToPlayIdx];
        Debug.Log($"Loading map: {mapToLoad}");
    
        isRacing = false;  // 로딩 시작 전에 false로 설정
        Loading.LoadScene(mapToLoad);
        raceToPlayIdx++;
    }

    //Customize Button Event
    public void GoCustomizeScene()
    {
        SceneManager.LoadScene("Customize");
    }

    public void SetMatchingStatus(MatchingUpdate mu)
    {
        matchingStatus = new Tuple<int, int>(mu.CurrentPlayers, mu.RequiredPlayers);
    }
    
    public Tuple<int,int> GetMatchingStatus()
    {
        return matchingStatus;
    }
    
}
