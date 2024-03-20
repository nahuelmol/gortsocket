import (
    "fmt"
    "log"
    "os/exec"
)

struct FFmpegCommand {
    action string
    cv string
    ca string
}

func (ff FFmpegCommand) Add(input, output, cv, cv) string {
    ff.action = "ffmpeg "
    ff.input = "-i " +input+ " "
    ff.output = output + " "
    ff.cv = "-c:v "cv + " "
    ff.ca = "-c:a "ca + " "

    dicommand = ff.action + ff.input + ff.cv + ff.ca + ff.output 
    fmt.Printf(dicommand)

    return dicommand 
}

func StartFfmpeg(){
    command, e := FFmpegCommand.Add("input.mp4", "output.mp4" ,"libx264", "aac")
    if e != {
        log.Println("error creating the command")
    }
    cmd := exec.Command("bash", "c" , command)

    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Println("err:", err)
        return
    }
    log.Println(string(out))

}
