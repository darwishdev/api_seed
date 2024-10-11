@echo off
setlocal enabledelayedexpansion
set count=1
for %%f in (*.webp) do (
    ren "%%f" rhss-premium_studio-!count!.webp
    set /a count+=1
)
