$script = $args[0]
$password = $args[1]
$salt = $args[2]

if (3 -gt $args.count){
	Write-Output "Decrypt and run .ps1 script encrypted using 'Out-EncryptedScript'"
	Write-Output "usage: run_encrypted_script.ps1 <script> <password> <salt>"
	Exit
}

$cmd = Get-Content $script 
Invoke-Expression $cmd
$decrypted = de $password $salt
Invoke-Expression $decrypted
