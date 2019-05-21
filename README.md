# Watch 
    A file system watcher that runs a bash script when a regex is matched on a
    file in a watched directory.

## Usage:
* run - `go run watch.go <directory> <regex> <script>`.
    * \<directory\> - The directory to watch over.
    * \<regex\> - Regular expression to match in order to run the script.
    * \<script\> - The file path to the script to execute when a file is changed. 
* build - `go build watch.go`.

## Notes: 
    The file name of the file that was modified is passed to the bash script as a
    command line argument and can be accessed with $1 inside of the script.
## Example:
    The command `go run watch.go . .c build.sh` would watch the current and any sub
    directory for changes in .c files and execute the build bash script if they change. 