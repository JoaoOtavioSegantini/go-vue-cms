import { mount } from "@vue/test-utils";
import EditPage from "@/components/Pages/Edit.vue";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
import { createStore } from "vuex";

describe("EditPage.vue", () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });

  window.axios = {
    get: jest
      .fn(() => {
        return {
          then: jest.fn(() => "your faked response"),
        };
      })
      .mockName("axiosget"),
    put: jest
      .fn(() => {
        return {
          then: jest.fn(() => "your faked response"),
        };
      })
      .mockName("axiosPut"),
  };

  const mockRoute = {
    params: {
      id: 1,
    },
  };
  const mockRouter = {
    push: jest.fn(),
  };
  const createVuexStore = () => {
    return createStore({
      state() {
        return {
          Pages: {
            onePage: {
              ID: "1",
              title: "Post 1",
              body: "A simple page text body",
              slug: "simple-slug-text",
            },
          },
        };
      },
      mutations: {
        updatePage(state) {
          state.onePage = {
            ID: "1",
            slug: "John Doe",
            title: "johndoe@gmail.com",
          };
        },
      },
      actions: {
        getPage(ctx, res) {
          ctx.commit("updatePage", res);
        },
        updatePage(ctx, data) {},
      },
    });
  };

  function factory() {
    const store = createVuexStore();

    return mount(EditPage, {
      mocks: {
        axios: window.axios,
      },
      global: {
        plugins: [store],
        mocks: {
          $router: mockRouter,
          $route: mockRoute,
        },
      },
      data() {
        return {
          editor: ClassicEditor,
        };
      },
    });
  }

  it("renders props.msg when passed", async () => {
    const msg = "AÇÕES Post 1 Edição de página do siteTítuloURLConteúdoSalvar";
    const wrapper = factory();

    wrapper.findAll("input")[0].setValue("acbdfg@gmail.com");
    wrapper.findAll("input")[1].setValue("987654");

    await wrapper.find("form").trigger("submit");
    await wrapper.find("button").trigger("click");
    await wrapper.vm.save();
    expect(mockRouter.push).toBeCalledWith({ path: "/pages/1" });
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
