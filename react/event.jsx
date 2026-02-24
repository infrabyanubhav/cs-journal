function Bulb(){
    let lightsOn=false;
    let OnState,offState;

    function toggleLight(){
        lightsOn=!lightsOn
        if (lightsOn){
                    OnState="On"
                    return OnState
        }else{
            offState="off"
                    return offState

        }

    }
    return(
        <>
        {lightsOn?OnState:offState}
        <button onClick={toggleLight}>
        </button>
        </>
    )
}

//Error Code 