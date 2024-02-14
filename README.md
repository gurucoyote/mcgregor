MyEvents Command Line Application
=================================

This simple command-line application is designed to process calendar events from `.ics` files within a specified directory. Here's what it does:

- **Find `.ics` Files**: It searches for any `.ics` (iCalendar) files in the given folder.
- **Parse Events**: It parses the found `.ics` files to extract event details.
- **Format Event Data**: For each event, it formats the start and end date, the summary (title), and the description (notes) into a single line.
- **Filter Past Events**: It filters out any events that have already occurred, only displaying future events.

Usage
-----
To use the application, provide the path to the directory containing `.ics` files as an argument:

```sh
myevents <path-to-directory>
```

The output will list the upcoming events in the following format:

```
start-date,weekday,summary,description,end-date
```

Where:
- `start-date` and `end-date` are in `YYYY-MM-DD` format.
- `weekday` is the day of the week when the event starts.
- `summary` is the title of the event.
- `description` is a brief note about the event, truncated to 100 characters if necessary.

Requirements
------------
The application requires the `github.com/arran4/golang-ical` package to parse `.ics` files. Make sure to have this dependency installed before running the application.

Contributing
------------
Contributions to this project are welcome. Please feel free to submit issues or pull requests for improvements or bug fixes.

License
-------
This project is open-source and available under the MIT License.
