open System

let consoleLock = obj ()

let log message =
    lock consoleLock (fun () -> printfn "%s" message)

// Each task represents work that spends time waiting for an external resource.
let performTask (taskName: string) (delayMs: int) =
    async {
        log (sprintf "%s started (delay: %d ms)" taskName delayMs)
        do! Async.Sleep delayMs
        log (sprintf "%s finished" taskName)
        return sprintf "%s result" taskName
    }

log "Starting three tasks concurrently..."

let tasks =
    [ performTask "Task 1" 1500
      performTask "Task 2" 500
      performTask "Task 3" 1000 ]

let results =
    tasks
    |> Async.Parallel
    |> Async.RunSynchronously

log "\nAll tasks completed. Results:"
results
|> Array.iter (fun result -> log (sprintf "- %s" result))