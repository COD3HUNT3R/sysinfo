// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

import "github.com/google/uuid"

// Product information.
type Product struct {
	Name    string    `json:"name,omitempty"`
	Vendor  string    `json:"vendor,omitempty"`
	Family  string    `json:"family,omitempty"`
	Version string    `json:"version,omitempty"`
	Serial  string    `json:"serial,omitempty"`
	UUID    uuid.UUID `json:"uuid,omitempty"`
	SKU     string    `json:"sku,omitempty"`
}

func (si *SysInfo) getProductInfo() {
	si.Product.Name = slurpFile("/sys/class/dmi/id/product_name")
	si.Product.Vendor = slurpFile("/sys/class/dmi/id/sys_vendor")
	si.Product.Family = slurpFile("/sys/class/dmi/id/product_family")
	si.Product.Version = slurpFile("/sys/class/dmi/id/product_version")
	si.Product.Serial = slurpFile("/sys/class/dmi/id/product_serial")
	si.Product.SKU = slurpFile("/sys/class/dmi/id/product_sku")

	uid, err := uuid.Parse(slurpFile("/sys/class/dmi/id/product_uuid"))
	if err == nil {
		si.Product.UUID = uid
	}

	// try a fallback to device-tree (ex: dmi is not available on ARM devices)
	// full details: https://www.devicetree.org/specifications/

	// on linux root path is /proc/device-tree (see: https://github.com/torvalds/linux/blob/v5.9/Documentation/ABI/testing/sysfs-firmware-ofw)
	if si.Product.Name == "" {
		si.Product.Name = slurpFile("/proc/device-tree/model")
	}

	if si.Product.Serial == "" {
		si.Product.Serial = slurpFile("/proc/device-tree/serial-number")
	}
}
