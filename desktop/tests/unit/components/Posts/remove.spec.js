import { mount } from "@vue/test-utils";
import RemovePage from "@/components/Posts/Remove.vue";
import { createStore } from "vuex";

describe("RemovePost.vue", () => {
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
    delete: jest
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
          Posts: {
            onePost: {
              ID: "1",
              title: "Post 1",
              body: "A simple post text body",
              slug: "simple-slug-text",
            },
          },
        };
      },
      mutations: {
        updatePost(state, res) {
          state.onePost = res;
        },
      },
      actions: {
        getPost(ctx, res) {
          ctx.commit("updatePost", res);
        },
        removePost(ctx, res) {},
      },
    });
  };

  function factory() {
    const store = createVuexStore();

    return mount(RemovePage, {
      global: {
        plugins: [store],
        mocks: {
          axios: window.axios,
          $route: mockRoute,
          $router: mockRouter,
        },
      },
    });
  }

  it("renders props.msg when passed", async () => {
    const msg =
    "AÇÕES Post 1 Remoção de artigo do blog Tem certeza que quer remover este artigo, essa ação não poderá ser desfeita! Não removerApagar definitivamente"
    const wrapper = factory();

    await wrapper.vm.remove();

    expect(wrapper.text()).toMatch(msg);

    await wrapper.findAll("a")[1].trigger("click");
    await wrapper.findAll("a")[0].trigger("click");
    console.log(wrapper.emitted());

    expect(wrapper.html()).toMatchSnapshot();
  });
});
