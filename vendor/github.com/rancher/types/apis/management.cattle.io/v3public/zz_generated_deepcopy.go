package v3public

import (
	reflect "reflect"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActiveDirectoryProvider).DeepCopyInto(out.(*ActiveDirectoryProvider))
			return nil
		}, InType: reflect.TypeOf(&ActiveDirectoryProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthProvider).DeepCopyInto(out.(*AuthProvider))
			return nil
		}, InType: reflect.TypeOf(&AuthProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthProviderList).DeepCopyInto(out.(*AuthProviderList))
			return nil
		}, InType: reflect.TypeOf(&AuthProviderList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AzureADLogin).DeepCopyInto(out.(*AzureADLogin))
			return nil
		}, InType: reflect.TypeOf(&AzureADLogin{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AzureADProvider).DeepCopyInto(out.(*AzureADProvider))
			return nil
		}, InType: reflect.TypeOf(&AzureADProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BasicLogin).DeepCopyInto(out.(*BasicLogin))
			return nil
		}, InType: reflect.TypeOf(&BasicLogin{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FreeIpaProvider).DeepCopyInto(out.(*FreeIpaProvider))
			return nil
		}, InType: reflect.TypeOf(&FreeIpaProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GenericLogin).DeepCopyInto(out.(*GenericLogin))
			return nil
		}, InType: reflect.TypeOf(&GenericLogin{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubLogin).DeepCopyInto(out.(*GithubLogin))
			return nil
		}, InType: reflect.TypeOf(&GithubLogin{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubProvider).DeepCopyInto(out.(*GithubProvider))
			return nil
		}, InType: reflect.TypeOf(&GithubProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LocalProvider).DeepCopyInto(out.(*LocalProvider))
			return nil
		}, InType: reflect.TypeOf(&LocalProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*OpenLdapProvider).DeepCopyInto(out.(*OpenLdapProvider))
			return nil
		}, InType: reflect.TypeOf(&OpenLdapProvider{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryProvider) DeepCopyInto(out *ActiveDirectoryProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryProvider.
func (in *ActiveDirectoryProvider) DeepCopy() *ActiveDirectoryProvider {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProvider) DeepCopyInto(out *AuthProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProvider.
func (in *AuthProvider) DeepCopy() *AuthProvider {
	if in == nil {
		return nil
	}
	out := new(AuthProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProviderList) DeepCopyInto(out *AuthProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProviderList.
func (in *AuthProviderList) DeepCopy() *AuthProviderList {
	if in == nil {
		return nil
	}
	out := new(AuthProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureADLogin) DeepCopyInto(out *AzureADLogin) {
	*out = *in
	out.GenericLogin = in.GenericLogin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureADLogin.
func (in *AzureADLogin) DeepCopy() *AzureADLogin {
	if in == nil {
		return nil
	}
	out := new(AzureADLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureADProvider) DeepCopyInto(out *AzureADProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureADProvider.
func (in *AzureADProvider) DeepCopy() *AzureADProvider {
	if in == nil {
		return nil
	}
	out := new(AzureADProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureADProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicLogin) DeepCopyInto(out *BasicLogin) {
	*out = *in
	out.GenericLogin = in.GenericLogin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicLogin.
func (in *BasicLogin) DeepCopy() *BasicLogin {
	if in == nil {
		return nil
	}
	out := new(BasicLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FreeIpaProvider) DeepCopyInto(out *FreeIpaProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FreeIpaProvider.
func (in *FreeIpaProvider) DeepCopy() *FreeIpaProvider {
	if in == nil {
		return nil
	}
	out := new(FreeIpaProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FreeIpaProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericLogin) DeepCopyInto(out *GenericLogin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericLogin.
func (in *GenericLogin) DeepCopy() *GenericLogin {
	if in == nil {
		return nil
	}
	out := new(GenericLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubLogin) DeepCopyInto(out *GithubLogin) {
	*out = *in
	out.GenericLogin = in.GenericLogin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubLogin.
func (in *GithubLogin) DeepCopy() *GithubLogin {
	if in == nil {
		return nil
	}
	out := new(GithubLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubProvider) DeepCopyInto(out *GithubProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubProvider.
func (in *GithubProvider) DeepCopy() *GithubProvider {
	if in == nil {
		return nil
	}
	out := new(GithubProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GithubProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalProvider) DeepCopyInto(out *LocalProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalProvider.
func (in *LocalProvider) DeepCopy() *LocalProvider {
	if in == nil {
		return nil
	}
	out := new(LocalProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLdapProvider) DeepCopyInto(out *OpenLdapProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLdapProvider.
func (in *OpenLdapProvider) DeepCopy() *OpenLdapProvider {
	if in == nil {
		return nil
	}
	out := new(OpenLdapProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenLdapProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
