package main

func main() {
	tv := &TV{}
	remoteControle := &RemoteControle{}

	remoteControle.ExecuteCommand(&TvOnCommand{tv: tv})
	remoteControle.ExecuteCommand(&TvOffCommand{tv: tv})
}
