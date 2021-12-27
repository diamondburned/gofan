# gofan

Small fan control program using curves, because fancontrol doesn't work.

## Usage

```sh
gofan \
	-function EaseInOutCubic -interval 2s \
	-pmin 10    -pmax 255   -p "/sys/devices/platform/asus-nb-wmi/hwmon/hwmon1/pwm1" \
	-tmin 58000 -tmax 85000 -t "/sys/devices/platform/coretemp.0/hwmon/hwmon5/temp3_input"
```

## Reference

https://easings.net/
