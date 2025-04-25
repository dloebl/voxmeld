# voxmeld
VoxMeld is a tool for The Elder Scrolls IV: Oblivion Remastered that restores original voice files and re-enables full voiced dialogue in additional languages, including German and French

> [!WARNING]
> There are a lot of manual steps required and it isn't (yet) straightforward to build a voice mod from the original voice files. At this point, building this manually is only intended for experienced users. I'll update this README once I improved/automated this process.

# Manual steps required to create the voice mod
1. Extract the original voice MP3s from the "Oblivion - Voices1.bsa" and "Oblivion - Voices2.bsa"
2. Convert the MP3s to WAVs with ffmpeg. Make sure the file names is something like "redguard_m_<action>.wav"
3. Import the WAVs into a new Audiokinetic Wwise project and convert them to WEMs with a codec of your choice (eg. Vorbis)
4. Unpack the "OblivionRemastered-Windows.pak" with UnrealPak.exe from the Unreal Engine 5
5. Copy the English(US) BNKs to a temporary folder (BNKs/)
8. Create the output folders: mkdir out_bnks/ && mkdir out_wems/
6. Run ./voxmeld <input.wem> on all replacement WEMs - this will patch the BNK, update the name of the WEM and save it to a new folder (out_bnks/ and out_wems/)
7. Repack everything in out_bnks/ and out_wems/ with UnrealPak.exe (important: keep the original folder structure from "OblivionRemastered-Windows.pak")
8. Place the new PAK in your ~mods/ folder in OblivionRemastered - the file name should be set to something like "german-voxmeld_P.pak"
9. Enjoy Oblivion Remastered in German or French!
