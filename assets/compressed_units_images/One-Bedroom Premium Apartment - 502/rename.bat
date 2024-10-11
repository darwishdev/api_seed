@echo off
setlocal enabledelayedexpansion
set count=0
for %%f in (*.webp) do (
    ren "%%f" rhss-1br_premium-!count!.webp
    set /a count = count+1
)
