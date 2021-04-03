package mmd

// AnimationHelperPoseOption is option for AnimationHelper's pose method.
type AnimationHelperPoseOption func(map[string]interface{}) error

// ResetPose is ...
func ResetPose(b bool) AnimationHelperPoseOption {
	return func(m map[string]interface{}) error {

		m["resetPose"] = b

		return nil
	}
}

// InverseKinematics is ...
func InverseKinematics(b bool) AnimationHelperPoseOption {
	return func(m map[string]interface{}) error {

		m["ik"] = b

		return nil
	}
}

// Grant is ...
func Grant(b bool) AnimationHelperPoseOption {
	return func(m map[string]interface{}) error {

		m["grant"] = b

		return nil
	}
}
