using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class OtherPlayer : MonoBehaviour
{
    public Vector3 destination;
    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        transform.position = destination;
    }
}
