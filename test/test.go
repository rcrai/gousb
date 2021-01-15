package main_test

// func TestNewController(t *testing.T) {
// 	var vendor, product uint16 = 0x2886, 0x0018

// 	var ctx = gousb.NewContext()
// 	ctx.Debug(1) // debug 4 could
// 	defer ctx.Close()

// 	// Iterate through available Devices, finding all that match a known VID/PID.
// 	vid, pid := ID(vendor), ID(product)

// 	devs, err := ctx.OpenDevices(filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var controllers []Controller
// 	for i := range devs {
// 		controllers = append(controllers, NewController(devs[i]))
// 	}
// 	return controllers, nil

// 	ctrls, err := FilterControllers(ctx, func(desc *DeviceDesc) bool {
// 		isValid := desc.Vendor == vid && desc.Product == pid

// 		fmt.Printf("list: Bus %d. Port %d. Addr %d ---- vendor: %v, product: %v : %v (%v)  ",
// 			desc.Bus, desc.Port, desc.Address,
// 			desc.Vendor, desc.Product, desc.String(), isValid,
// 		)

// 		// log.Infof("vendor: %v, product: %v : %v (%v)  -- bus:%d,port:%d,addr:%d", desc.Vendor, desc.Product, desc.String(), isValid,
// 		// 	desc.Bus, desc.Port, desc.Address,
// 		// )
// 		// this function is called for every device present.
// 		// Returning true means the device should be opened.
// 		if isValid {
// 			fmt.Printf("Got Device: vendor: %v, product: %v : %v", desc.Vendor, desc.Product, desc.String())
// 		}
// 		return isValid
// 	})
// }
