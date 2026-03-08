module main

go 1.25.3

replace handler => ./handler

replace userHandler => ./userHandler

require handler v0.0.0-00010101000000-000000000000

require userHandler v0.0.0-00010101000000-000000000000 // indirect
