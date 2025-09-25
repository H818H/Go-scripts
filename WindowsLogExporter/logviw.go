package main

import ("fmt"
		"os"
		"os/exec"
		"bufio"
		"time")

func get_log_info(command string) (string, error){
	proc := exec.Command("powershell.exe")
	
	stdin, err := proc.StdinPipe()
	if err != nil{
		return "", fmt.Errorf("* error creating StdinPipe: %w", err)
	}
	
	stdout, err := proc.StdoutPipe()
	if err != nil{
		return "", fmt.Errorf("* error creating StdoutPipe: %w", err)
	}
	
	if err := proc.Start(); err != nil{
		return "", fmt.Errorf("* error while creating powershell process: %w", err)
	}
	fmt.Println("process pid: ", proc.Process.Pid)
	
	_, err = stdin.Write([]byte(command + "\n"))
	if err != nil{
		return "", fmt.Errorf("* error trying to write command on powershell: %w", err)
	}
	stdin.Close()
	
	scanner := bufio.NewScanner(stdout)
	var output string
	for scanner.Scan(){
		output += scanner.Text() + "\n"
	}
	
	if err := scanner.Err(); err != nil{
		return "", fmt.Errorf("* error trying to read output from powershell: %w", err)
	}
	
	if err := proc.Wait(); err != nil{
		return "", fmt.Errorf("* error while waiting for process: %w", err)
	}
	
	return output, nil
}

func main(){
	os.Mkdir("logs", 0755)
	
	var commandList = [4]string{
	"Get-WinEvent -LogName Application -MaxEvents 50|Select @{Name='TimeCreated';Expression={$_.TimeCreated.ToString('yyyy-MM-dd HH:mm:ss')}},EventId,Source,Message|ConvertTo-Json -Depth 4 > logs/application.json",
	"Get-WinEvent -LogName Security -MaxEvents 50|Select @{Name='TimeCreated';Expression={$_.TimeCreated.ToString('yyyy-MM-dd HH:mm:ss')}},EventId,Source,Message|ConvertTo-Json -Depth 4 > logs/security.json",
	"Get-WinEvent -LogName System -MaxEvents 50|Select @{Name='TimeCreated';Expression={$_.TimeCreated.ToString('yyyy-MM-dd HH:mm:ss')}},Id,ProviderName,Message|ConvertTo-Json -Depth 4 > logs/system.json",
	"Get-WinEvent -LogName Setup -MaxEvents 50|Select @{Name='TimeCreated';Expression={$_.TimeCreated.ToString('yyyy-MM-dd HH:mm:ss')}},Id,ProviderName,Message|ConvertTo-Json -Depth 4 > logs/setup.json"}
	
	for i := 0; i < len(commandList); i++{
		fmt.Printf("\n[%d]- creating/updating json file...\n", i+1)
	
		_, err := get_log_info(commandList[i])
		if err != nil{
			fmt.Println("* error while creating json: ", err)
		}else{
			fmt.Println("- created/updated without errors")
		}
	}
	
	time.Sleep(1 * time.Second)
}