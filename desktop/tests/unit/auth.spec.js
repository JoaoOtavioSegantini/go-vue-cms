import Auth from "@/components/Auth.vue";
import { mount } from "@vue/test-utils";

// const sinon = require('sinon');

// const mockRequest = () => {
//   return {
//     users: []
//   };
// };

// const mockResponse = () => {
//   const res = {};
//   res.status = sinon.stub().returns(res);
//   res.json = sinon.stub().returns(res);
//   return res;
// };

window.axios = {
  post: jest
    .fn(() => {
      return {
        then: jest.fn(() => "your faked response"),
      };
    })
    .mockName("axiosPost"),
};

describe("Auth.vue", () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });

  function factory() {
    return mount(Auth, {
      mocks: {
        axios: window.axios,
      },
      data() {
        return {
          user: {
            email: "acbdfg@gmail.com",
            password: "987654",
          },
        };
      },
    });
  }

  it("renders props when passed", async () => {
    const msg = "Autenticação Informe suas credenciaisEmailSenhaAcessar";

    const wrapper = factory();

    console.log(wrapper.vm);
    wrapper.findAll("input")[0].setValue("acbdfg@gmail.com");
    wrapper.findAll("input")[1].setValue("987654");

    await wrapper.find("form").trigger("submit");
    console.log(wrapper.emitted());

    await wrapper.vm.auth();

    expect(window.axios.post).toBeCalledWith(
      "/api/token",
      { email: "acbdfg@gmail.com", password: "987654" },
      {
        swalMessage: "Você acessou com sucesso",
        swalTitle: "Autenticado com sucesso",
      }
    );
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
