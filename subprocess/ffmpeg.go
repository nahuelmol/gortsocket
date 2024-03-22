package subprocess


import (
    "log"
    "os/exec"
)

type FFmpegCommand struct {
    action string
    cv string
    ca string
    raw string

    input string
    output string
}

type Check struct {
    target string
    raw string
}

func Start(target string) *Check {
    ctarget := target
    //I must build the raw command
    craw := target + " "
    return &Check {
        target:ctarget,
        raw:craw,
    }
}

func NewCommand(input, output, cv, ca string) *FFmpegCommand {
    FFaction := "ffmpeg "
    FFinput := "-i " +input+ " "
    FFoutput := output + " "
    FFcv := "-c:v " + cv + " "
    FFca := "-c:a " + ca + " "

    FFraw := FFaction + FFinput + FFcv + FFca + FFoutput 
    return &FFmpegCommand {
        action:FFaction,
        input:FFinput,
        output:FFoutput,
        cv:FFcv,
        ca:FFca,
        raw:FFraw,
    }
}

func ChecktheRunning(proc string){
    //checking that videos are runnig
    command := Start(proc)
    
    cmd := exec.Command("bash", "c" , command.raw)
        
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Println("err:", err)
        return
    }
    log.Println(string(out))
}


func StartFfmpeg(){
    command := NewCommand("input.mp4", "output.mp4" ,"libx264", "aac")
    
    cmd := exec.Command("bash", "c" , command.raw)

    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Println("err:", err)
        return
    }
    log.Println(string(out))
}
