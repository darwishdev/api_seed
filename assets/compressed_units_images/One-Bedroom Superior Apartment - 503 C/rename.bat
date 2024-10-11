@echo off
setlocal enabledelayedexpansion
set count=1
for %%f in (*.webp) do (
    ren "%%f" rhss-1br_superior-!count!.webp
    set /a count = count+1
)
