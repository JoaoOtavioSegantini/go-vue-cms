import Auth from "@/components/Auth.vue";
import { mount } from "@vue/test-utils";

describe("Auth.vue", () => {
  function factory() {
    return mount(Auth, {
      // data() {
      //   return {
      //     user: {
      //       email: "acbdfg@gmail.com",
      //       password: "987654"
      //     },
      //   };
      // },
    });
  }

  it("renders props when passed", async () => {
    const msg = "Autenticação Informe suas credenciaisEmailSenhaAcessar";

    const wrapper = factory();
    console.log(wrapper.vm);
    // expect(wrapper.vm.auth()).toHaveBeenCalled();
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
