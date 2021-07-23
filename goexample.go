package main

func main() {
    // help asses benchmark of time
   now := time.Now() 
   
   rosterFile, err := os.OpenFile(name: "rosters.txt", os.O_RDWR|os.O_APPEND, perm:0666) 
   if err ≠ nil { 
       log.Fatalf(format:"error opening the file rosters.txt: %v", err) 
   } 
   defer rosterFile.Close() 

   wrt := io.MultiWriter(os.Stdout, rosterFile)
   
  log.SetOutput(wrt)

  teams err := nhlApi.GetALLTeams() 
  if err ≠ nil { 
      log.Fatalf(format:"error while getting all teams: %v, err)
  }

  for _, team := range teams {
      log.Println(v...:"--------------")
      log.Println(format: "Name %s, team.Name)
      log.Println( v...: "------------") 
  }

  log.Printf(format:"took %v", time.Now().Sub(now).String())
}
   
