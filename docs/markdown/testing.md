# Testing

## Make setup files
- `$ cd scripts` (if you haven't already)
- `$ sh buildTest.sh`

## Create a testing VM
NOTE: For VMs i use [virtual box](https://www.virtualbox.org/)
- Create a VM and install windows 10 (*64 bit*)
- Add the `build` folder as shared folder (don't add the root of this repo because that will cause an issue where it reads the wrong config file ([issue](https://github.com/dennis1248/Automated-Windows-10-configuration/issues/10))) 
- Create a soft link from the `build` network folder to the desktop
- Activate windows (if windows is not activated some things registry things might not work)
- Make a snapshot of the machine (make sure the machine is idle because if not windows might crash later)
- close the VM

## Running a test on the VM
- Create a test build shown at the top how to
- Restore the VM to the created snapshot and start the VM
- Run the software
- Check if everything works
- Close the VM

## Tips
- Install chocolatey and create a new snapshot so you don't have to wait every time for chocolatey to install
- Don't make a snapshot while the build folder is open because network locations don't automaticly refresh
